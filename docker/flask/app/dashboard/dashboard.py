from flask import request, session, render_template, redirect, url_for, current_app
import random, json
from .. import lobbies, socketio, mutex
from ..models import User, Settings
from . import dashBP

@dashBP.route("/dashboard")
def dashboard():
    user_id = session.get('user_id')

    user_settings = Settings.query.filter_by(user_id=session["user_id"]).first()
    settings = {"port":"50051", "window":"firefox", "host":session["username"]}
    if user_settings:
        settings = json.loads(user_settings.settings_json)

    current_app.logger.info(f"{settings}")


    if not user_id:
        return redirect(url_for("login.login"))
    
    return render_template("dashboard.html", user_id=user_id)

# @socketio.on("create lobby")
@dashBP.route("/create_lobby")
def create_lobby():
    lobby_id = '%010x' % random.randrange(16**10)
    while lobby_id in lobbies:
        lobby_id = '%010x' % random.randrange(16**10)
    
    user_settings = Settings.query.filter_by(user_id=session["user_id"]).first()
    settings = {"port":"50051", "window":"firefox", "host":session["username"]}
    if user_settings:
        settings = json.loads(user_settings.settings_json)
    
    with mutex:
        lobbies[lobby_id] = {"owner":{"owner_id":session.get("user_id"), "owner_host":settings["host"], "owner_username":session.get("username")},
                            "participants":[]
                            }
    return redirect(url_for("lobby.view_lobby", lobby_id=lobby_id))