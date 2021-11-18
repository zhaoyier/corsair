import * as React from "react";

const styles = require("./index.scss");

interface PageUrl {
	name: string;
	hash: string;
}

export interface HeaderProps {
	module: string;
	parent?: PageUrl;
	current?: PageUrl;
	cb?: () => void;
}

export const Header = (props: HeaderProps) => {
	return (
		<section className={styles.header}>
			<section className={styles.crumbs}>
				{props.module}&nbsp;/&nbsp;
				{props.parent ? (
					<span
						className={styles.parent}
						onClick={() => (window.location.hash = props.parent.hash)}>
						{props.parent.name}&nbsp;/&nbsp;
					</span>
				) : null}
				{props.current ? (
					<span
						className={styles.current}
						onClick={() => (window.location.hash = props.current.hash)}>
						{props.current.name}
					</span>
				) : null}
			</section>
		</section>
	);
};
