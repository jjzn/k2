const link = document.querySelector(`nav a[href="${location.pathname}"]`);

// TODO: do something when we're on /
if (link) { // Means we are on a special site, such as current month
    link.classList.add('current');
    document.title = `${link.textContent} — k2`;
}
