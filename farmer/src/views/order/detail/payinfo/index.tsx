import * as React from "react";
import { PayDetail } from "./payDetail";
import { PaySurvey } from "./paySurvey";
export function PayInfo() {
	return (
		<div style={{ marginTop: 30 }}>
			<h3>支付明细</h3>
			<PayDetail />
			<PaySurvey />
		</div>
	);
}
