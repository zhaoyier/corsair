
import App from "../views/order";
import Mount from "../components/Mount";

Mount(App, undefined, undefined);
if (module.hot) {
	module.hot.accept("../views/order", () => {
		Mount(App, undefined, undefined);
	});
}
