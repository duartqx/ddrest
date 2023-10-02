function storeOnLocalStorage(key, values) {
  localStorage.setItem(key, JSON.stringify(values));
}

function loadFromLocalStorage(key) {
  let values = JSON.parse(localStorage.getItem(key));
  if (!values) {
    return false;
  }
  for (let [key, value] of Object.entries(values)) {
    if (key.includes("::")) {
      [_, key] = key.split("::");
    }
    let inp = htmx.find(`#${key}`);
    inp.value = value.toString();
  }
  return true
}

function resetRealEstateForm() {
  htmx.find("#realEstateFilterForm").reset();
  localStorage.removeItem("realEstateFilterFormCache");
  htmx.trigger("#realEstateFilterForm", "submit");
}
