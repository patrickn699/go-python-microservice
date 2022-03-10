from re import S
from flask import render_template, request, redirect, url_for, Flask, jsonify
import requests



app = Flask(__name__)

TEMPLATE_DIRS = ['./templates']
STATIC_DIRS = ['./static']

@app.route('/', methods=['GET', 'POST'])
def index():
    return render_template('index.html')


@app.route('/fil', methods=['POST'])
def fil():
    if request.method == 'POST':
        file = request.files["file"]
        file.save('./saved/' + file.filename)
        send_file('./saved/' + file.filename)
        return redirect(url_for('index'))


def send_file(filename):
    url = "http://localhost:9000/convert"
    files = {'file': open(filename, 'rb')}
    r = requests.post(url, files=files)
    print("made a post request to conversion service:",r.text)






if __name__ == '__main__':
    app.run(debug=True,port=5000)