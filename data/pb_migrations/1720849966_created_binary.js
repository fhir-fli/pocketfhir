/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "4fg7jtycbsqx95o",
    "created": "2024-07-13 05:52:46.097Z",
    "updated": "2024-07-13 05:52:46.097Z",
    "name": "binary",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ocyegtov",
        "name": "versionId",
        "type": "number",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "noDecimal": false
        }
      },
      {
        "system": false,
        "id": "0bhqpa2x",
        "name": "resource",
        "type": "json",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "maxSize": 5242880
        }
      }
    ],
    "indexes": [],
    "listRule": "@request.auth.id != ''",
    "viewRule": "@request.auth.id != ''",
    "createRule": "@request.auth.id != ''",
    "updateRule": "@request.auth.id != ''",
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("4fg7jtycbsqx95o");

  return dao.deleteCollection(collection);
})