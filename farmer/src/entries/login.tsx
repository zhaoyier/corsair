
import App from "../views/login";
import Mount from "../components/Mount";

Mount(App, true, true);
if (module.hot) {
	module.hot.accept("../views/login", () => {
		Mount(App, true, true);
	});
}
