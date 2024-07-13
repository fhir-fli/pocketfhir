/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "z9gdtx7jay3bo74",
    "created": "2024-07-13 05:52:46.994Z",
    "updated": "2024-07-13 05:52:46.994Z",
    "name": "organizationhistory",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "wra7heib",
        "name": "fhirId",
        "type": "text",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "fhjoxhax",
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
        "id": "yevqv70l",
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
  const collection = dao.findCollectionByNameOrId("z9gdtx7jay3bo74");

  return dao.deleteCollection(collection);
})