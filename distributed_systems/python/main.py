from flask import Flask, jsonify, request
import threading
import time

app = Flask(__name__)

data_store = {}

@app.route('/data', methods=['POST'])
def add_data():
    data = request.get_json()
    data_store[data['key']] = data['value']
    return jsonify({'message': 'Data added successfully!'}), 200

@app.route('/data/<key>', methods=['GET'])
def get_data(key):
    value = data_store.get(key, 'Key not found!')
    return jsonify({'value': value}), 200

def background_task():
    while True:
        print("Running background task...")
        time.sleep(10)

if __name__ == '__main__':
    threading.Thread(target=background_task).start()
    app.run(debug=True, host='0.0.0.0', port=5000)
