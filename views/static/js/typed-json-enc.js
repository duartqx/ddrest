htmx.defineExtension("typed-json-enc", {
  onEvent: function (name, evt) {
    if (name === "htmx:configRequest") {
      evt.detail.headers["Content-Type"] = "application/json";
    }
  },

  encodeParameters: function (xhr, parameters, elt) {
    let parsedValues = new Object();
    xhr.overrideMimeType("text/json");
    for (const [typedKey, value] of Object.entries(parameters)) {
      let [type, key] = typedKey.split("::");

      switch (type) {
        case "String":
          parsedValues[key] = value;
          break;
        case "Int":
          parsedValues[key] = parseInt(value);
          break;
        case "Float":
          parsedValues[key] = parseFloat(value);
          break;
        case "Bool":
          parsedValues[key] = value.toLowerCase() === "true";
          break;
        default:
          parsedValues[key] = value;
      }
    }
    return JSON.stringify(parsedValues);
  },
});
