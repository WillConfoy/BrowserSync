from flask import Blueprint

dashBP = Blueprint("dashboard", __name__,
                   template_folder="templates")

from . import dashboard