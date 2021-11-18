import * as React from "react";
import cs from "classnames";
import { Layout, Menu } from "antd";
const styles = require("./index.scss");
import { MenuItem, fetchMenu } from "utils/menu";
import { saveClick } from "./util";
import TopBar from "components/TopBar";
import { getUserInfo } from "utils/user";
import { redirectToLogin } from "utils/auth";
import { TUser } from "genServices/AdminLoginService";
import { Link } from "react-router-dom";

interface SidebarProps {
	collapsed: boolean;
}

const { Sider } = Layout;
const { SubMenu } = Menu;

const initUserInfo: TUser = {
	username: "",
	systems: []
};

const titleName = (name: string, zhName: string, enName: string, icon?: string): any => {
	const isZh = navigator.language === "zh-CN";

	return (
		<span>
			{icon && <i className={cs(styles.icons, icon)} />}
			<span className={styles.title}> {(isZh ? zhName : enName) || name}</span>
		</span>
	);
};

const Sidebar = (props: SidebarProps) => {
	const { collapsed } = props;
	const [openKeys, setOpenKeys] = React.useState([]);
	const [selectedKeys, setSelectedKeys] = React.useState([]);
	const [menuList, setMenuList] = React.useState([]);
	const [userInfo, setUserInfo] = React.useState(initUserInfo);

	React.useEffect(() => {
		const newUserInfo = getUserInfo();
		if (!userInfo) {
			const { hash, pathname } = window.location;
			redirectToLogin(pathname + hash);
			return;
		}
		setUserInfo(newUserInfo);

		fetchMenu().then(data => {
			checkMenu(data);
		});
	}, []);

	const checkMenu = (menuData: MenuItem[]) => {
		const menuList = menuData;
		const { openKeys, selectedKeys } = getOpenKeysAndSelectedKeys(menuList);
		setOpenKeys(openKeys);
		setSelectedKeys(selectedKeys);
		setMenuList(menuList);
	};

	const filterControlledSystems = (controlledSystems: string[], menu: MenuItem[]) => {
		return menu.reduce<MenuItem[]>((pValue, menuItem) => {
			const newMenuItem = { ...menuItem };
			if (Array.isArray(menuItem.submenu)) {
				newMenuItem.submenu = filterControlledSystems(controlledSystems, menuItem.submenu);
			}
			if (
				(Array.isArray(newMenuItem.submenu) && newMenuItem.submenu.length > 0) ||
				controlledSystems.includes(menuItem.name)
			) {
				pValue.push(newMenuItem);
			}
			return pValue;
		}, []);
	};

	const getOpenKeysAndSelectedKeys = (menuList: MenuItem[], groupNames: string[] = []) => {
		const currentURL = new URL(window.location.href);
		return menuList.reduce<{ openKeys: string[]; selectedKeys: string[] }>(
			(pValue, cValue) => {
				const hasSubMenu = Array.isArray(cValue.submenu) && cValue.submenu.length > 0;
				if (cValue.href && !hasSubMenu && equalMenuItem(cValue.href, currentURL)) {
					pValue.selectedKeys.push(cValue.href);
					pValue.openKeys = pValue.openKeys.concat(groupNames);
				}
				if (hasSubMenu) {
					const { openKeys, selectedKeys } = getOpenKeysAndSelectedKeys(
						cValue.submenu,
						groupNames.concat(cValue.name)
					);
					pValue.openKeys = pValue.openKeys.concat(openKeys);
					pValue.selectedKeys = pValue.selectedKeys.concat(selectedKeys);
				}
				return pValue;
			},
			{ openKeys: [], selectedKeys: [] }
		);
	};

	const equalMenuItem = (href: string, url: URL) => {
		if (typeof href === "string") {
			let checkURL: URL = null;
			try {
				checkURL = new URL(
					`${window.location.protocol}//${window.location.host}${
						href.startsWith("/") ? "" : "/"
					}${href}`
				);
			} catch (e) {
				return false;
			}
			const hashResult = checkURL.hash === "" || checkURL.hash === url.hash;
			let paramsResult = true;
			for (const pair of checkURL.searchParams.entries()) {
				if (url.searchParams.get(pair[0]) !== pair[1]) {
					paramsResult = false;
				}
			}
			return checkURL.pathname === url.pathname && paramsResult && hashResult;
		}
		return false;
	};

	const clickMenuItem = (eve, href: string, item: MenuItem) => {
		eve.preventDefault();
		const ele = saveClick(item.name, href); // 当前选中菜单项  存在session里面
		// window.open(ele.link, "_blank"); //打开新标签
		window.location.href = ele.link;   //不打开新标签
		// <Link to={`${ele.link}`}>{ele.link}</Link>
		return false;
	};

	const renderMenu = (data: MenuItem[]) => {
		return data.map((item, i) => {
			const { submenu, zhName, enName, name, href, icon } = item;
			if (submenu && submenu.length > 0) {
				const cls: string[] = [];
				if (openKeys.includes(name)) {
					cls.push(styles.menuOpened);
				} else {
					cls.push(styles.menuHover);
				}
				return renderMenu(submenu).every(item => item === null) ? null : (
					<SubMenu
						key={name}
						title={titleName(name, zhName, enName, icon)}
						className={cls.join(" ")}>
						{renderMenu(submenu)}
					</SubMenu>
				);
			} else {
				let nameHtml = titleName(name, zhName, enName, icon);
				return (
					<Menu.Item key={`${item.isHistory ? "hot-" : ""}${href ? href : i}`}>
						<a
							className={cs(styles.menuItem, {
								[styles.isHot]: item.isHistory
							})}
							href={href}
							onClick={e => clickMenuItem(e, href, item)}>
							{nameHtml}
						</a>
					</Menu.Item>
				);
			}
		});
	};

	return (
		<Sider trigger={null} collapsible collapsed={collapsed} collapsedWidth={50} width={150} theme="light">
			<div className={styles.sideMenuBar}>
				<TopBar collapsed={collapsed} />
				<div className={styles.menuContainerDiv}>
					<Menu
						selectedKeys={selectedKeys}
						openKeys={openKeys}
						onOpenChange={openKeys => setOpenKeys(openKeys as any)}
						onClick={selectedKey => setSelectedKeys([selectedKey.key])}
						className={styles.menuContainer}>
						{renderMenu([...menuList])}
					</Menu>
				</div>
			</div>
		</Sider>
	);
};

export default Sidebar;
