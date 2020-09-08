#!/usr/bin/env python3.6
# -*- coding: utf-8 -*- 
"""
server.py
Python 3.6 or newer required.
"""

import json
import os
import uuid

import sqlite3
from flask import Flask, abort , jsonify, request

app = Flask(__name__)

@app.route('/create',  methods=['POST'])
def create_user():
    try:
        if not request.json or not 'user' in request.json:
            abort(400)
        _id = uuid.uuid1()
        print(request.json['user'])
        db = sqlite3.connect('demo.db')
        cursor = db.cursor()
        cursor.execute('''INSERT INTO demo(user, tok, stats)
                    VALUES(?,?,?)''', (request.json['user'], str(_id), str("open")))
        db.commit()
        return jsonify({'id':_id}), 200
    except Exception as e:
        return jsonify(error=str(e)), 403

if __name__ == '__main__':
    app.run(debug=True)