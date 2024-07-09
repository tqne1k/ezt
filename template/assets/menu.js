document.addEventListener("DOMContentLoaded", function() {
    var menuHtml = `
        <ul>
            <li><a href="/relation_ui">Relations</a></li>
            <li><a href="/add_relation_ui">Add Relation</a></li>
            <li><a href="/object_ui">Objects</a></li>
            <li><a href="/add_object_ui">Add Object</a></li>
        </ul>
    `;
    document.getElementById("menu-container").innerHTML = menuHtml;
});