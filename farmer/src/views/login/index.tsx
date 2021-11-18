import * as React from "react";
import cs from "classnames";
import LoginForm from "./loginForm";
const styles = require("./index.scss");

class Login extends React.Component<{}, {}> {
	render() {
		return [
			<div className={styles.mask} key={1}>
				<p>
					<span>EZShip Operation System</span>
				</p>
			</div>,
			<section className={styles.box} key={2}>
				<div className={cs(styles.logo)}>
					<span className={styles.logoName}>Ezship</span>
					<span>OMS后台</span>
				</div>
				<LoginForm />
			</section>
		];
	}
}

export default Login;
