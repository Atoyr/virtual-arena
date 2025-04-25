# API エンドポイント

HTTP メソッド | パス | 説明 | パラメータ／Body
GET    | /maps              | 利用可能なマップ一覧を取得 | クエリ: ?page=&limit=
POST   | /maps              | 新しいマップを登録（TMX/JSON＋メタ情報） | Body(JSON): { name, description, data: <tmj/json>, tilesets: [...] }
GET    | /maps/:mapId       | マップの JSON（TMJ）データ本体 | URL: :mapId
PUT    | /maps/:mapId       | マップ情報の更新 | Body(JSON): { name?, description?, data?, tilesets? }
DELETE | /maps/:mapId       | マップの削除 | URL: :mapId
GET    | /maps/:mapId/tilesets                  | マップに紐付くタイルセット一覧を取得 | クエリ: ?page=&limit=
GET    | /maps/:mapId/tilesets/:tilesetName.png | 指定タイルセット画像 | URL: :mapId, :tilesetName.png


