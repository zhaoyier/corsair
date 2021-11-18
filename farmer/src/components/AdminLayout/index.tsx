import * as React from "react";
import cs from "classnames";
import { Layout } from "antd";
import "./index.scss";
import Sidebar from "./Sidebar";
import LayoutHeader from "./layoutHeader";
import { getUserInfo } from "utils/user";
import { redirectToLogin, loginTimeout } from "utils/auth";
import { checkExp } from "utils/time";
import { isUAT } from "utils/url";
const styles = require("./index.scss");

const { Content } = Layout;

export interface AdminLayoutProps {
	/** 是否隐藏菜单 */
	hideMenu: boolean;
	hidePadding: boolean;
}

interface AdminLayoutState {
	collapsed: boolean; // 是否折叠菜单
}

export default class AdminLayout extends React.Component<AdminLayoutProps, AdminLayoutState> {
	constructor(props) {
		super(props);
		this.state = {
			collapsed: false
		};
	}
	componentDidMount() {
		const userInfo = getUserInfo();
		if (!userInfo && !window.location.pathname.match(/login/)) {
			redirectToLogin(window.location.href);
		} else {
			if (!window.location.pathname.match(/login/) && checkExp(userInfo.timeToExpire)) {
				loginTimeout();
			}
		}
		if (userInfo) {
			(window as any).evaluationInit &&
				(window as any).evaluationInit({
					username: userInfo.name, // passport用户名，必传，建议在登录/获取用户信息之后，调用evaluationInit方法
					appName: "ezship-oms", // appName，必传, 接入passport的配置
					accessKey: "ezship-oms", // accessKey，必传, 接入passport的配置
					isOnline: isUAT() ? false : true // 是否生产环境，默认false, 需要业务系统自行判断是否线上环境
				});
		}
	}

	toggleCollapsed = () => {
		this.setState({
			collapsed: !this.state.collapsed
		});
	};

	render() {
		const { collapsed } = this.state;
		const { hideMenu } = this.props;
		if (hideMenu) {
			return <Layout className={styles.adminGrid}>{this.props.children}</Layout>;
		}
		return (
			<Layout className={styles.adminGrid}>
				<Sidebar collapsed={collapsed} />
				<Layout className={styles["site-layout"]}>
					<LayoutHeader collapsed={collapsed} toggleCollapsed={this.toggleCollapsed} />
					<Content className={cs(styles["site-layout-background"], styles.content)}>
						{this.props.children}
					</Content>
				</Layout>
			</Layout>
		);
	}
}
