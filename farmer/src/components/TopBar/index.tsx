import * as React from "react";
import cs from "classnames";
import { getUserInfo } from "utils/user";

const styles = require("./index.scss");

interface TopBarProps {
	collapsed?: boolean;
	className?: string;
	isMainMenu?: boolean;
}

const TopBar = (props: TopBarProps) => {
	let { isMainMenu } = props;
	const { className, collapsed } = props;
	const [userInfo, setUserInfo] = React.useState(getUserInfo());
	const [show, setShow] = React.useState(false);
	isMainMenu = Boolean(isMainMenu) ? isMainMenu : true;

	React.useEffect(() => {
		const userInfos = getUserInfo();
		setUserInfo(userInfos);
	}, []);

	const handleLogOut = () => {
		window.location.href = "/login.html";
	};
	const formateName = (name: string) => {
		const regx = /[a-z][A-Z]/gi;
		if (regx.test(name)) {
			let nextName = name.slice(0, 2);
			const firstLetter = name.slice(0, 1);
			nextName = nextName.replace(firstLetter, firstLetter.toUpperCase());
			return nextName;
		}
		return name.slice(0, 1);
	};
	const switchLogOut = (status: boolean) => {
		if (status !== show) {
			setShow(status);
		}
	};

	return (
		<section className={cs(styles.topBar, className ? className : "")}>
			<div className={cs(styles.logo)}>
				<span className={styles.logoName}>{!collapsed ? "Ezship" : "E"}</span>
				{!collapsed && <span>OMS</span>}
			</div>
			{!collapsed && (
				<div
					className={cs(styles.userInfo, show ? styles.active : "")}
					onMouseMove={() => switchLogOut(true)}
					onMouseOut={() => switchLogOut(false)}>
					<span className={styles.name}>
						{userInfo && formateName(userInfo.username)}
					</span>
					<a onClick={handleLogOut}>
						<i className={styles.icon} />
						Log Out
					</a>
				</div>
			)}
		</section>
	);
};

export default TopBar;
