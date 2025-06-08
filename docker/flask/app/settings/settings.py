from flask import request, jsonify, session, flash, redirect, url_for, render_template
import json
from ..models import User, Settings
from .. import db
from . import settingsBP



@settingsBP.route('/user-settings', methods=['GET', 'POST'])
def get_user_settings():
    # Make sure user is logged in and accessing their own settings
    user_id = session["user_id"]
    user = User.query.get_or_404(user_id)

    if not user:
        flash("Error: User not found")
        return redirect(url_for("login.register"))

    # Load existing settings or create default
    user_settings = Settings.query.filter_by(user_id=user_id).first()
    if user_settings:
        settings_data = user_settings.settings_json

    if request.method == 'POST':
        # Expecting form data with JSON string or individual fields
        # For simplicity, let's assume a textarea called 'settings'
        new_settings = request.form.get('settings')
        try:
            # Validate JSON format
            parsed_settings = json.loads(new_settings)
        except json.JSONDecodeError:
            flash("Invalid JSON format.", "danger")
            return render_template('settings.html', settings=new_settings)

        if user_settings:
            user_settings.settings_json = new_settings
        else:
            user_settings = Settings(user_id=user_id, settings_json=new_settings)
            db.session.add(user_settings)

        db.session.commit()
        flash("Settings updated successfully.", "success")
        return redirect(url_for('dashboard.dashboard'))

    return render_template('settings.html', settings=settings_data)