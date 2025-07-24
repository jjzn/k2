const form = document.forms[0];
const location_field = form['location'];
const list = document.querySelector('#location-suggestions');

const MAX_SUGGESTIONS = 3;
const TYPING_THRESHOLD_MS = 500;

let typing_timeout;

// TODO: update version
const headers = new Headers({
    'User-Agent': 'k2/0.2'
});

const attribution = document.createElement('p');
attribution.innerHTML = 'Data from <a href="https://openstreetmap.org/copyright">OpenStreetMap</a>';

async function load_suggestions() {
    const query = encodeURIComponent(location_field.value);
    // TODO: cache results in the server
    const resp = await fetch(`https://nominatim.openstreetmap.org/search?q=${query}&limit=${MAX_SUGGESTIONS}&format=jsonv2`, { headers });
    const data = await resp.json();

    list.innerHTML = ''; // clear the list

    // TODO: constuct custom name based on the address fields
    for (const place of data) {
        const item = document.createElement('li');
        const link = document.createElement('a');

        link.textContent = place.display_name;
        link.href = 'javascript:';
        link.addEventListener('click', () => {
            location_field.value = place.display_name;
            list.innerHTML = ''; // clear the suggestions list on click
        });

        item.append(link);
        list.append(item);
        list.append(attribution);
    }

    list.hidden = false;
    location_field.after(list);
}

location_field.addEventListener('input', async () => {
    // Prevent location suggestions while still typing, we don't want to
    // overload the server with requests
    clearTimeout(typing_timeout);
    typing_timeout = setTimeout(() => load_suggestions(), TYPING_THRESHOLD_MS);

    list.innerHTML = ''; // clear the list
});
