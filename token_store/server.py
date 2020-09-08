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
def create():
    try:       
        if request.method == 'POST' and not request.json or not 'user' in request.json:
            abort(400)
        else:
            _id = uuid.uuid1()
            user_name = request.json['user']
            db = sqlite3.connect('demo.db')
            cursor = db.cursor()
            cursor.execute('''INSERT INTO demo(user, id, stats)
                           VALUES(?,?,?)''', (user_name, str(_id), str("open")))
            db.commit()
            return jsonify({'id':_id, 'message': '{} has been created successfully'.format(user_name)}), 200
    except Exception as e:
        return jsonify(error=str(e)), 403

@app.route('/update',  methods=['POST'])
def update():
    try:       
        if request.method == 'POST' and not request.json or not 'old' in request.json or not 'new' in request.json:
            abort(400)
        else:
            db = sqlite3.connect('demo.db')
            cursor = db.cursor()
            old = request.json['old']
            new = request.json['new']
            cursor.execute('update demo set user = "{0}" where user = "{1}";'.format(new, old))
            cur = db.cursor()
            cur.execute('select id from demo where user = "{0}";'.format(new))
            ids = cur.fetchone()
            db.commit()
            return jsonify({'id':ids[0],'message': '{0} has been updated to {1} successfully'.format(old, new)}), 200
            
    except Exception as e:
        return jsonify(error=str(e)), 403

@app.route('/delete',  methods=['POST'])
def delete():
    try:       
        if request.method == 'POST' and not request.json or not 'user' in request.json:
            abort(400)
        else:
            user_name = request.json['user']
            db = sqlite3.connect('demo.db')
            cursor = db.cursor()
            cursor.execute('delete from demo where user = "{0}"'.format(user_name))
            db.commit()
            return jsonify({'message': '{} has been deleted successfully'.format(user_name)}), 200
            
    except Exception as e:
        return jsonify(error=str(e)), 403

@app.route('/user',  methods=['POST'])
def user():
    try:       
        if request.method == 'POST' and not request.json or not 'user' in request.json:
            abort(400)
        else:
            _id = uuid.uuid1()
            user_name = request.json['user']
            db = sqlite3.connect('demo.db')
            cursor = db.cursor()
            cursor.execute('select * from demo where user = "{0}"'.format(user_name))
            records = cursor.fetchall()
            user = records[0][0]
            _id = records[0][1]
            stats = records[0][2]
            db.commit()
            return jsonify({'user': user, 'id':_id, 'stats':stats}), 200
            
    except Exception as e:
        return jsonify(error=str(e)), 403
    
if __name__ == '__main__':
    app.run(debug=True)