from flask import request, session, render_template, redirect, flash, url_for
import json
from .. import db
from ..models import User, Settings
from . import loginBP


@loginBP.route('/register', methods=['GET', 'POST'])
def register():
    if session.get("user_id"):
            redirect(url_for("dashboard.dashboard"))
    if request.method == "GET":
        return render_template("register.html")
    username = request.form["username"]
    password = request.form["password"]

    if User.query.filter_by(username=username).first():
        flash("Error: Username already exists", "danger")
        return render_template("register.html")

    if len(username) > 120:
        flash("Error: Username too long", "warning")
        return render_template("register.html")

    user = User(username=username)
    user.set_password(password)

    db.session.add(user)
    db.session.commit()

    default_settings = json.dumps({"port":"50051", "window":"chrome", "host":username, "display":"1"})
    user_settings = Settings(user_id=user.id, settings_json=default_settings)

    db.session.add(user_settings)
    db.session.commit()

    session["user_id"] = user.id
    session["username"] = username
    return redirect(url_for("dashboard.dashboard"))


@loginBP.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == "POST":
        if session.get("user_id"):
            redirect(url_for("dashboard.dashboard"))

        username = request.form["username"]
        password = request.form["password"]

        user = User.query.filter_by(username=username).first()
        if user and user.check_password(password):
            session["user_id"] = user.id
            session["username"] = username
            return redirect(url_for("dashboard.dashboard"))
        else:
            flash("Error: Invalid username or password", "danger")
            return render_template("login.html")

    return render_template("login.html")

@loginBP.route("/logout")
def logout():
    session.clear()
    return redirect(url_for("login.login"))