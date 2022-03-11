import { useState, useEffect, useCallback } from 'react';
import type { FC } from 'react';
import {
  Card,
  CardHeader,
  Divider,
  Table,
} from '@mui/material';
// import { customerApi } from '../../../api/customer-api';
// import type { CustomerLog } from '../../../types/customer';
import { MoreMenu } from '../../more-menu';
import { Scrollbar } from '../../scrollbar';

export const CustomerLogs: FC = (props) => {

  return (
    <Card {...props}>
      <CardHeader
        action={<MoreMenu />}
        title="Recent Logs"
      />
      <Divider />
      <Scrollbar>
        <Table sx={{ minWidth: 700 }}>
        </Table>
      </Scrollbar>
    </Card>
  );
};
