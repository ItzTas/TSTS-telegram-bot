from flask import Flask, request, jsonify
from typing import Dict, Tuple

app: Flask = Flask(__name__)


@app.route("/chat", methods=["GET"])
def chat() -> Tuple[Dict[str, str], int]:
    try:
        data = request.get_json()
        content = data["content"]
        return jsonify({"content_is": content}), 200
    except KeyError:
        return jsonify({"error": "missing 'content' body paramether"}), 400
    except Exception as e:
        return jsonify({"error": f"unexpected error \n {e}"}), 500
