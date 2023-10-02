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
}

function lazyLoad(url, target, source) {
  let values = parseTypedKeys(htmx.values(htmx.find(source)));
  fetch(url, {
    method: "POST",
    body: JSON.stringify(values),
    headers: { "Content-Type": "application/json" },
  })
    .then((res) => res.text())
    .then((res) => {
      htmx.find(target).innerHTML = res;
    });
}

function lazyLoadRealEstate() {
  lazyLoad("/filterEstate", "#result", "#realEstateFilterForm");
}

function resetRealEstateForm() {
  htmx.find("#realEstateFilterForm").reset();
  localStorage.removeItem("realEstateFilterFormCache");
  lazyLoadRealEstate();
}
