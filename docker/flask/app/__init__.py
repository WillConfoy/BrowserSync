import threading
import os
from flask import Flask
from flask_socketio import SocketIO
from flask_sqlalchemy import SQLAlchemy
# from werkzeug.middleware.proxy_fix import ProxyFix

socketio = SocketIO()
db = SQLAlchemy()
lobbies = {}
mutex = threading.Lock()

def create_app(debug=False):
    """Create an application."""
    app = Flask(__name__)
    app.debug = debug
    app.config['SECRET_KEY'] = os.getenv("FLASK_SECRET_KEY")
    app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:sys@localhost:5432/postgres'
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    # app.wsgi_app = ProxyFix(app.wsgi_app, x_for=1, x_proto=1, x_host=1, x_prefix=1)

    db.init_app(app)
    socketio.init_app(app)

    from .models import User, Settings

    from .index import indexBP
    app.register_blueprint(indexBP)

    from .settings import settingsBP
    app.register_blueprint(settingsBP)

    from .login import loginBP
    app.register_blueprint(loginBP)

    from .dashboard import dashBP
    app.register_blueprint(dashBP)

    from .lobbies import lobbyBP
    app.register_blueprint(lobbyBP)

    # Create tables if they don't exist
    with app.app_context():
        db.create_all()

    return app