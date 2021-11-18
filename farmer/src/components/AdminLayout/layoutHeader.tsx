import * as React from "react";
import { Layout, Tooltip, message, Badge } from "antd";
import cs from "classnames";
import { MenuUnfoldOutlined, MenuFoldOutlined, LogoutOutlined } from "@ant-design/icons";
import { redirectToLogin } from "utils/auth";
import { removeUserInfo, getUserInfo } from "utils/user";

interface LayoutHeaderProps {
	noticeInfo?: {
		announcementCount: number;
		mailCount: number;
	};
	collapsed: boolean; // 菜单是收起
	toggleCollapsed: () => void;
}
const { Header } = Layout;
const styles = require("./index.scss");

const getMessageCount = (count: number): "99+" | number => {
	return count > 99 ? "99+" : count;
};

const handleLogOut = () => {
	message.success("退出登录成功！");
	removeUserInfo();
	redirectToLogin("/index.html");
};

const LayoutHeader = (props: LayoutHeaderProps) => {
	const { collapsed, toggleCollapsed, noticeInfo } = props;
	const userInfo = getUserInfo();
	const total = noticeInfo ? noticeInfo.announcementCount + noticeInfo.mailCount : 0;

	return (
		<Header className={cs(styles["site-layout-background"], styles.layoutHeader)}>
			{React.createElement(collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
				className: styles.trigger,
				onClick: toggleCollapsed
			})}
			<div className={cs(styles.userInfo, styles.active)}>
				<div className={styles.layoutHeaderWrap}>
					<div className="welcome">
						<span>
							<Badge
								offset={[-3, 8]}
								count={getMessageCount(total)}
								className={styles.noticeWrap}>
								<img src={require("./image/notice.svg")} />
							</Badge>
						</span>
						<img
							style={{ width: "30px", margin: "0 10px 0 20px" }}
							src={require("./image/Default head image.svg")}
						/>
						<span>{userInfo && userInfo.username}</span>
						<span className={styles.dividing} />
					</div>
				</div>
				<Tooltip title="Log Out" placement="bottomLeft">
					<a onClick={handleLogOut}>
						<LogoutOutlined />
					</a>
				</Tooltip>
			</div>
		</Header>
	);
};

export default LayoutHeader;
