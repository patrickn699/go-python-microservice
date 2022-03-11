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
        height = request.form["height"]
        width = request.form["width"]
        percent = request.form["percent"]
        water = request.form["watermark"]
        file.save('./saved/' + file.filename)
        send_file('./saved/' + file.filename,height,width,percent,water)
        return redirect(url_for('index'))


def send_file(filename,height,width,percent,water):
    url = "http://localhost:9000/convert"
    files = {'file': open(filename, 'rb')}
    try:
        r = requests.post(url, files=files,data={'height':height,'width':width,'percent':percent,'watermark':water})
        print("made a post request to conversion service:",r.text)
    except Exception as e:
        return render_template('error.html', error=e)






if __name__ == '__main__':
    app.run(debug=True,port=5000)