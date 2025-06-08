from flask import Blueprint

lobbyBP = Blueprint("lobby", __name__,
                    template_folder="templates")

from . import lobby