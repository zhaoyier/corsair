import * as React from "react";
import * as cs from "classnames";

const styles = require("./index.scss");

export interface LayoutProps {
	header: React.ReactNode;
	showBar?: boolean;
	children?: any;
	nothasBottom?: boolean;
	style?: any;
}

const CommonLayout = (props: LayoutProps) => {
	const { showBar, header, children, nothasBottom, style } = props;
	return (
		<section className={styles.layoutTotal}>
			<div className={nothasBottom ? styles.commonHeader : styles.layoutHeader}>{header}</div>
			<div
				style={style}
				className={cs(styles.layoutContent, !showBar ? styles.notShowBar : "")}>
				{children}
			</div>
		</section>
	);
};

export default CommonLayout;
