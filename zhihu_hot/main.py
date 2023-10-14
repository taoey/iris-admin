import sqlite3
import json
import time
import requests

dbName = "data/zhihu_hot.db"

def saveResp(json_data):
    conn = sqlite3.connect(dbName)
    cursor = conn.cursor()
    creation_date = int(time.time())  # 当前时间戳
    # 创建新表（如果不存在），用于存储原始 JSON 数据
    cursor.execute('''CREATE TABLE IF NOT EXISTS zhihu_hot_json (
                    id INTEGER PRIMARY KEY,
                    json_data TEXT,
                    create_time INTEGER)''')

    # 插入原始 JSON 数据
    cursor.execute("INSERT INTO zhihu_hot_json (json_data,create_time) VALUES (?,?)", (json.dumps(json_data,ensure_ascii=False),creation_date))
    conn.commit()
    conn.close()

def saveData(json_data):
    # 解析 JSON 数据
    json_data_all = json.loads(json_data)
    conn = sqlite3.connect(dbName)
    cursor = conn.cursor()
    creation_date = int(time.time())  # 当前时间戳
    
    for data in json_data_all['data']:
        # 提取所需字段
        question_id = data['target']['id']
        title = data['target']['title']
        hot_score = float(data['detail_text'].replace("万热度", ""))
        url = data['target']['url'].replace("api", "www").replace("questions", "question")
        excerpt = data['target']['excerpt']

        # 创建表（如果不存在）
        cursor.execute('''CREATE TABLE IF NOT EXISTS zhihu_hot (
                        id INTEGER PRIMARY KEY,
                        question_id TEXT,
                        title TEXT,
                        url TEXT,
                        hot_score DOUBLE,
                        excerpt TEXT,
                        create_time INTEGER)''')

        # 插入数据，包括创建日期
        cursor.execute("INSERT INTO zhihu_hot (question_id,title,url,hot_score, excerpt, create_time) VALUES (?,?,?,?, ?, ?)",
                       (question_id, title, url,hot_score, excerpt, creation_date))
        conn.commit()

    # 关闭数据库连接
    conn.close()


def getZhuHotData():
    url = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=100&desktop=false"
    response = requests.request("GET", url)
    return response.text


def main():
    resp = getZhuHotData()
    saveResp(resp)
    saveData(resp)
    return
main()
