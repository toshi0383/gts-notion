import Foundation

guard let authToken: String = ProcessInfo.processInfo.environment["NOTION_AUTH_TOKEN"] else {
    fatalError("NOTION_AUTH_TOKEN is required")
}

struct NotionRecord: Decodable {
    let courseCode: String
    let tire: Int
    let fuel: Int
    let category: String
    let fav: String
    let laps: Int
    let date: String
    let tires: [TireName]
    let series: Series
}

enum Series: Int, Decodable {
    case manufacturer, nations

    var description: String {
        switch self {
        case .manufacturer:
            return "マニュファクチャラーズ"
        case .nations:
            return "ネイションズ"
        }
    }
}

enum TireName: Int, Decodable, CustomStringConvertible {
    case confortHard = 0, confortMedium, confortSoft, confortSuperSoft
    case sportHard, sportMedium, sportSoft, sportSuperSoft
    case racingHard, racingMedium, racingSoft, racingSuperSoft
    case racingInter, racingWet
    case dirt, spike

    var description: String {
        switch self {
        case .confortHard: return "コンフォート・ハード"
        case .confortMedium: return "コンフォート・ミディアム"
        case .confortSoft: return "コンフォート・ソフト"
        case .confortSuperSoft: return "コンフォート・スーパーソフト"
        case .sportHard: return "スポーツ・ハード"
        case .sportMedium: return "スポーツ・ミディアム"
        case .sportSoft: return "スポーツ・ソフト"
        case .sportSuperSoft: return "スポーツ・スーパーソフト"
        case .racingHard: return "RH"
        case .racingMedium: return "RM"
        case .racingSoft: return "RS"
        case .racingSuperSoft: return "RSS"
        case .racingInter: return "IM"
        case .racingWet: return "HW"
        case .dirt: return "ダートタイヤ"
        case .spike: return "雪用スパイクタイヤ"
        }
    }
}

// NOTE: Not used
enum PropertyKey: String {
	// タイヤ消耗倍率
	case tire = "\"Yj\\" // number
	// シリーズ
	case series = ")KB)" // select
	// 練習ステータス
	case status = ",n7O" // select
	// 燃料消耗倍率
	case fuel = "8Jfz" // number
	// カテゴリ、車種
	case category = ";}4]" // select
	// お気に入り
	case fav = "AAXJ" // checkbox
	// 戦略
	case strategy = "AD6a" // text
	// 周回
	case laps = "OgNk" // number
	// Date
	case date = "e;Kd" // date
	// 使用タイヤ
	case tires = "yrHc" // multi_select
	// Name
	case title = "title" // title
}

func request(url: URL, body: Data) -> URLRequest {
    var req = URLRequest(url: url)
    req.httpMethod = "POST"
    req.setValue("token_v2=\(authToken)", forHTTPHeaderField: "cookie")
    req.setValue("application/json", forHTTPHeaderField: "content-type")
    req.httpBody = body
    return req
}

@discardableResult
func logResponse(_ data: Data?, _ res: URLResponse?, _ err: Error?) -> Bool {
    var success = true
    if let err = err {
        print("\(err)")
        success = false
    }
    if let code = (res as? HTTPURLResponse)?.statusCode, !(200..<300).contains(code) {
        print("\(code)")
        success = false
    }
    if let data = data {
        print(String(data: data, encoding: .utf8)!)
    }
    return success
}

// process

print("Ctrl + C to finish or abort.")

let str: String = {
    var s = [String]()
    while let l = readLine() {
        s.append(l)
    }
    return s.joined()
}()

let data: Data = str.data(using: .utf8)!

let decoder = JSONDecoder()
decoder.keyDecodingStrategy = .convertFromSnakeCase
let notionRecords: [NotionRecord] = try! decoder.decode([NotionRecord].self, from: data)

for (i, r) in notionRecords.enumerated() {

    let uuid: String = UUID().uuidString.lowercased()
    let nowUnixtime: Int = Int(Date().timeIntervalSince1970 * 1000)

    let tiresStr: String = r.tires.map({ "\"\($0)\"" }).joined(separator: ",")

    let submitBody: Data = """
    {
        "operations": [{
            "id": "\(uuid)",
            "table": "block",
            "path": [],
            "command": "set",
            "args": {
                "type": "page",
                "id": "\(uuid)",
                "version": 1
            }
        }, {
            "id": "\(uuid)",
            "table": "block",
            "path": ["properties"],
            "command": "update",
            "args": {
                "e;Kd": [
                    ["‣", [
                        ["d", {
                            "type": "date",
                            "start_date": "\(r.date)"
                        }]
                    ]]
                ]
            }
        }, {
            "id": "\(uuid)",
            "table": "block",
            "path": [],
            "command": "update",
            "args": {
                "parent_id": "a16d13b7-f929-4c2e-8ae0-6fc7491f0435",
                "parent_table": "collection",
                "alive": true
            }
        }, {
            "table": "block",
            "id": "\(uuid)",
            "path": ["created_by"],
            "command": "set",
            "args": "4a3ea399-4382-497d-803d-e03ac2c89ac5"
        }, {
            "table": "block",
            "id": "\(uuid)",
            "path": ["created_time"],
            "command": "set",
            "args": \(nowUnixtime)
        }, {
            "table": "block",
            "id": "\(uuid)",
            "path": ["last_edited_time"],
            "command": "set",
            "args": \(nowUnixtime)
        }, {
            "table": "block",
            "id": "\(uuid)",
            "path": ["last_edited_by"],
            "command": "set",
            "args": "4a3ea399-4382-497d-803d-e03ac2c89ac5"
        }, {
            "table": "block",
            "id": "\(uuid)",
            "path": ["properties", "title"],
            "command": "set",
            "args": [
                ["\(courseNameMap[r.courseCode]!)"]
            ]
        },
        {"id":"\(uuid)","table":"block","path":["properties",";}4]"],"command":"set","args":[["\(r.category)"]]},
        {"id":"\(uuid)","table":"block","path":["properties","\\\"Yj\\\\"],"command":"set","args":[["\(r.tire)"]]},
        {"id":"\(uuid)","table":"block","path":["properties","8Jfz"],"command":"set","args":[["\(r.fuel)"]]},
        {"id":"\(uuid)","table":"block","path":["properties","yrHc"],"command":"set","args":[[\(tiresStr)]]},
        {"id":"\(uuid)","table":"block","path":["properties",",n7O"],"command":"set","args":[["検討中"]]},
        {"id":"\(uuid)","table":"block","path":["properties",")KB)"],"command":"set","args":[["\(r.series)"]]},
        {"id":"\(uuid)","table":"block","path":["properties","AAXJ"],"command":"set","args":[["No"]]}
        ]
    }
    """.data(using: .utf8)!

    let url = URL(string: "https://www.notion.so/api/v3/submitTransaction")!

    let session = URLSession.shared

    session.dataTask(with: request(url: url, body: submitBody)) { (data, res, err) in
        let success = logResponse(data, res, err)
        print("finished \(i + 1) / \(notionRecords.count), success: \(success)")
    }.resume()
}

dispatchMain()
