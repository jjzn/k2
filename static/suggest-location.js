const form = document.forms[0];
const location_field = form['location'];

const MAX_SUGGESTIONS = 3;

// TODO: provide attribution
// TODO: determine wether to listen to input or blur events
location_field.addEventListener('input', async () => {
    const query = encodeURIComponent(location_field.value);
    const resp = await fetch(`https://nominatim.openstreetmap.org/search?q=${query}&limit=${MAX_SUGGESTIONS}&format=jsonv2`); // TODO: set the user agent
    const data = await resp.json();

    const list = document.querySelector('#location-suggestions');
    list.innerHTML = ''; // clear the list

    // TODO: constuct custom name based on the address fields
    for (const place of data) {
        const item = document.createElement('li');
        const link = document.createElement('a');

        link.textContent = place.display_name;
        link.href = 'javascript:';
        link.addEventListener('click', () => {
            location_field.value = place.display_name;
        });

        item.append(link);
        list.append(item);
    }

    list.hidden = false;
    location_field.after(list);
});
