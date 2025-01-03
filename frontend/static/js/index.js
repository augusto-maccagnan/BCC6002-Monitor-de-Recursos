import Dashboard from "./views/Dashboard.js";
import Settings from "./views/Settings.js";

const navigateTo = (url) => {
    history.pushState(null, null, url);
    router();
};

const router = async () => {
    const routes = [{ path: "/", view: Dashboard }, { path: "/settings", view: Settings }, ,];
    const potentialMatches = routes.map((route) => {
        return {
            route: route,
            isMatch: location.pathname === route.path,
        };
    });

    let match = potentialMatches.find((potentialMatch) => potentialMatch.isMatch);

    if (!match) {
        match = {
            route: routes[0],
            isMatch: true,
        };
    }

    const view = new match.route.view();
    document.querySelector("#app").innerHTML = await view.getHtml();
    const formTempoAtualizacao = document.getElementById("formTempoAtualizacao");
    formTempoAtualizacao?.addEventListener("change", (e) => {
        e.preventDefault();
        const timer = new FormData(formTempoAtualizacao);
        const obj = Object.fromEntries(timer);
        const json = JSON.stringify(obj);
        localStorage.setItem("timer", json);
    });
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", (e) => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });
    router();
});
