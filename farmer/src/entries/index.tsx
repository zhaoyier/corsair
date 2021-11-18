
import App from "../views/index";
import Mount from "../components/Mount";

Mount(App, undefined, undefined);
if (module.hot) {
	module.hot.accept("../views/index", () => {
		Mount(App, undefined, undefined);
	});
}
