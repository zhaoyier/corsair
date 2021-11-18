import * as React from "react";
import Search from "./search";
import CreateOrder from "./createOrder";
import OrderTable from "./orderTable";
import CommonLayout from "components/CommonLayout";
import { Header } from "components/CommonHeader";
import { listHeader } from "../constant";

export default function List() {
	return (
		<CommonLayout style={{ padding: 0 }} header={<Header {...listHeader} />} showBar={false}>
			<Search />
			<CommonLayout
				style={{ margin: 0, padding: 0 }}
				header={<CreateOrder />}
				showBar={false}>
				<OrderTable />
			</CommonLayout>
		</CommonLayout>
	);
}
