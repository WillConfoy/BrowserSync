from flask import redirect, session, url_for
from . import indexBP

@indexBP.route('/')
def index():
    check = session.get("user_id")
    if not check:
        return redirect(url_for("login.login"))
    return redirect(url_for("dashboard.dashboard"))