from flask import Blueprint

indexBP = Blueprint("index", __name__,
                      template_folder="templates")

from . import index