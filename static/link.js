const elem = document.getElementById('desc');

const url = /((https?:\/\/)?[\w\-~]+(\.[\w\-~]+)+(\/[\w\-\.~]*)*(#[\w\-\.~]*)?)|(\w+:.+)/g;

const link = function(match) {
    const href = /^\w+:/.test(match)
        ? match : 'http://' + match;

    return `<a href="${href}">${match}</a>`;
};

elem.innerHTML = elem.innerHTML.replace(url, link);
