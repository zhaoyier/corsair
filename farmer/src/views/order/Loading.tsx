import * as React from "react";
import { Col, Spin, Row } from "antd";
export const Loading = ({ isLoading, error }) => {
	// Handle the loading state
	if (isLoading) {
		return (
			<Row>
				<Col
					span={12}
					offset={12}
					style={{ height: "500px", lineHeight: "60px", paddingTop: "60px" }}>
					<Spin tip="Loading..." />
				</Col>
			</Row>
		);
	}
	// Handle the error state
	else if (error) {
		return <div>Sorry, there was a problem loading the page.</div>;
	} else {
		return null;
	}
};
