## azs-json-filter-golang

### A simple script for filtering a json file.

#### Disclaimer:
The script is not universal, it was needed for a single use of a large amount of data.

Developed for filtering by companies of fuel stations for integration into Maps.

Filters data by ```properties.hintContent``` in the json structure as:

```json
[
  {
    "properties": {
      "hintContent": "someData",
      "balloonContentFooter": "someData",
      "balloonContentHeader": "someData",
      "iconCaption": "someData"
    },
    "geometry": {
      "coordinates": ["someData", "someData"],
      "type": "someData"
    },
    "options": {
      "preset": "someData",
      "iconColor": "someData"
    },
    "id": "someData",
    "type": "someData"
  },
  ...
]
```

