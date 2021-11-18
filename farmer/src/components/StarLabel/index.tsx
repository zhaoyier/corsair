import * as React from "react";

export function StarLabel(label: string | number) {
	return (
		<React.Fragment>
			<span
				style={{
					marginRight: "4px",
					color: "#ff4d4f",
					fontSize: "14px",
					fontFamily: "SimSun,sans-serif",
					lineHeight: 1
				}}>
				*
			</span>
			{label}
		</React.Fragment>
	);
}
