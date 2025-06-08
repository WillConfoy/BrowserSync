from flask import Blueprint

loginBP = Blueprint("login", __name__,
                    template_folder="templates")

from . import login