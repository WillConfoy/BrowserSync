from flask import  render_template, request, redirect, url_for, session, flash
from flask_socketio import leave_room, join_room, emit, send
import grpc, json
from .. import lobbies, socketio, mutex, db
from ..models import User, Settings
from . import lobbyBP
from .python_gather import gather_pb2, gather_pb2_grpc

allowTransfer = True

def emit_participants(lobby_id):
    lobby = lobbies[lobby_id]
    participants = [i["username"] for i in lobby["participants"]]
    emit("participants", participants, room=lobby_id)


@socketio.on("join")
def handle_join():
    lobby_id = session["lobby_id"]
    lobby = lobbies[lobby_id]

    user_settings = Settings.query.filter_by(user_id=session["user_id"]).first()
    settings = {"port":"50051", "window":"firefox"}
    if user_settings:
        settings = json.loads(user_settings.settings_json)

    with mutex:
        lobby["participants"].append({
            "user_id":session["user_id"],
            "host":settings["host"],
            "username":session["username"]})

    join_room(lobby_id)
    emit("message", f"{session["username"]} has joined the lobby", room=lobby_id)
    emit_participants(lobby_id)
    emit("allow_transfer_update", allowTransfer, room=lobby_id)

@socketio.on("disconnect")
def handle_disconnect():
    lobby_id = session["lobby_id"]
    lobby = lobbies[lobby_id]
    deleted = False

    with mutex:
        lobby["participants"] = [i for i in lobby["participants"] if i["user_id"] != session["user_id"]]
        if len(lobby["participants"]) == 0: del lobby
        deleted = True
    emit("message", f"{session["username"]} left the lobby", room=lobby_id)
    if not deleted: emit_participants(lobby_id)
    leave_room(lobby_id)

############## FORMAT ##############
# lobbies[lobby_id] = {"owner":{"owner_id":session.get("user_id"), "owner_remote_addr":request.remote_addr, "owner_username":session.get("username")},
#                     "participants":[]
#                     }
#
# AND
#
# lobby["participants"].append({
#             "user_id":session["user_id"],
#             "remote_addr":request.remote_addr,
#             "username":session["username"]})
####################################

@socketio.on("start lobby")
def handle_start():
    lobby = lobbies[session["lobby_id"]]
    addrstring = ""

    for machine in lobby["participants"]:
        addrstring += machine["host"]+"|"
    addrstring = addrstring[:-1]

    for machine in lobby["participants"]:
        user = User.query.get_or_404(machine["user_id"])
        if not user:
            print("ERROR GETTING USER!!!!")
            continue

        user_settings = Settings.query.filter_by(user_id=machine["user_id"]).first()
        settings = {"port":"50051", "window":"firefox"}
        if user_settings:
            settings = json.loads(user_settings.settings_json)

        is_owner = lobby["owner"]["owner_id"] == machine["user_id"]
        with grpc.insecure_channel(machine["host"]+':20601') as channel:
            stub = gather_pb2_grpc.GatherServiceStub(channel)
            machineInfoRequest = gather_pb2.MachineInfoRequest(port=settings["port"], ip=machine["host"], window=settings["window"])
            response = stub.SendMachineInfo(machineInfoRequest)
            if not response.success:
                print(f"FAILED SENDING MACHINE INFO TO {machine}")
            stateInfoRequest = gather_pb2.StateInfoRequest(leader=is_owner, allowtransfer=allowTransfer, addrstring=addrstring)
            response = stub.SendStateInfo(stateInfoRequest)
            if not response.success:
                print(f"FAILED SENDING STATE INFO TO {machine}")

@socketio.on("leave lobby")
def handle_leave():
    emit("left")

@socketio.on("allow_transfer_changed")
def handle_allow(allow_transfer):
    global allowTransfer
    allowTransfer = allow_transfer
    emit("allow_transfer_update", allowTransfer, room=session["lobby_id"])


# Dashboard should have create lobby button that sends you to your own lobby, as well as
# a join lobby field that lets you join a lobby by inputting the lobby ID

@lobbyBP.route('/lobby/<string:lobby_id>')
def view_lobby(lobby_id):
    lobby = lobbies.get(lobby_id)
    if not lobby:
        flash("Lobby not found.", "danger")
        return redirect(url_for('dashboard.dashboard'))
    
    session["lobby_id"] = lobby_id
    is_owner = lobby["owner"]["owner_id"] == session["user_id"]
    return render_template('lobby.html', lobby_id=lobby_id, is_owner=is_owner)
