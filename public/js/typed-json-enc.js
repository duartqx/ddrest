function parseTypedKeys(parameters) {
  let parsedValues = new Object();
  for (const [typedKey, value] of Object.entries(parameters)) {
    if (!value) {
      continue;
    } else if (!typedKey.includes("::")) {
      parsedValues[typedKey] = value;
      continue;
    }

    let [type, key] = typedKey.split("::");

    switch (type) {
      case "Int":
        parsedValues[key] = parseInt(value);
        break;
      case "Float":
        parsedValues[key] = parseFloat(value);
        break;
      case "Bool":
        parsedValues[key] =
          value.toLowerCase() === "true" || value.toLowerCase() === "1";
        break;
      default:
        parsedValues[key] = value;
    }
  }
  return parsedValues
}
htmx.defineExtension("typed-json-enc", {
  onEvent: function (name, evt) {
    if (name === "htmx:configRequest") {
      evt.detail.headers["Content-Type"] = "application/json";
    }
  },

  encodeParameters: function (xhr, parameters, elt) {
    xhr.overrideMimeType("text/json");
    let parsedValues = parseTypedKeys(parameters);
    return JSON.stringify(parsedValues);
  },
});
