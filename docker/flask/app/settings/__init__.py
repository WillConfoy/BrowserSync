from flask import Blueprint

settingsBP = Blueprint("settings", __name__,
                      template_folder="templates")

from . import settings