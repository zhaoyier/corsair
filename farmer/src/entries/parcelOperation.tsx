
import App from "../views/parcelOperation";
import Mount from "../components/Mount";

Mount(App, undefined, undefined);
if (module.hot) {
	module.hot.accept("../views/parcelOperation", () => {
		Mount(App, undefined, undefined);
	});
}
