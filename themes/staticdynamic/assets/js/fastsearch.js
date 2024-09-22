var fuse; // holds our search engine
var firstRun = true; // allow us to delay loading json data unless search activated
var list = document.getElementById('searchResults'); // targets the <ul>
var resultsAvailable = false; // Did we get any search results?

// ==========================================
// The event listener that waits for the input search box to be on focus
// 
//
document.getElementById("searchInput").addEventListener("focus", function() {
  //if(sessionStorage.getItem("firstrun") === null)
  //  sessionStorage.setItem("firstrun", "true"); // for the first time
  //if(sessionStorage.getItem("firstrun") == "true") {
  if (firstRun) {
    loadSearch(); // loads our json data and builds fuse.js search index
    firstRun = false;
  }  
  //  sessionStorage.setItem("firstrun", "false"); // let's never do this again
}); 


// ==========================================
// execute search as each character is typed
//
document.getElementById("searchInput").onkeyup = function(e) { 
  executeSearch(this.value);
}


// ==========================================
// fetch some json without jquery
//
function fetchJSONFile(path, callback) {
  var httpRequest = new XMLHttpRequest();
  httpRequest.onreadystatechange = function() {
    if (httpRequest.readyState === 4) {
      if (httpRequest.status === 200) {
        var data = JSON.parse(httpRequest.responseText);
          if (callback) callback(data);
      }
    }
  };
  httpRequest.open('GET', path);
  httpRequest.send(); 
}


// ==========================================
// load our search index, only executed once
// when the search box gets the focus
//
function loadSearch() { 
  fetchJSONFile('/index.json', function(data){

    var options = { // fuse.js options; check fuse.js website for details
      shouldSort: true,
      findAllMatches: true,
      location: 0,
      distance: 100,
      threshold: 0.4,
      minMatchCharLength: 3,
      keys: [
        {name:"title",weight:8},
        {name:"description",weight:6},
        {name:"tags",weight:6},
        {name:"summary",weight:5},
        {name:"contents", weight:1}
        ]
    };
    fuse = new Fuse(data, options); // build the index from the json file
  });
}


// ==========================================
// using the index we loaded when the search box gets the focus 
// a search query (for "term") every time a letter is typed
// in the search box
//
function executeSearch(term) {
  let results = fuse.search(term); // the actual query being run using fuse.js
  let searchitems = ''; // our results bucket

  if (results.length === 0) { // no results based on what was typed into the input box
    resultsAvailable = false;
    searchitems = '';
  } else { // build our html 
    for (let item in results.slice(0,5)) { // only show first 5 results
      searchitems += '<li><a href="' + results[item].item.permalink + '" tabindex="0">' + '<span class="results-title">' + results[item].item.title + '</span><br /><span class="results-description">'+ results[item].item.description +'</span></a></li>';
    }
    resultsAvailable = true;
  }
  document.getElementById("searchResults").innerHTML = searchitems;
}