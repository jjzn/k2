const link = document.querySelector(`nav a[href="${location.pathname}"]`);
link.classList.add('current');

const text = location.pathname === '/'
    ? 'Todas las entradas' : link.textContent;

document.title = `${text} — k2`;
