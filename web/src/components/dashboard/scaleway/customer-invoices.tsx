import { useState, useEffect, useCallback } from 'react';
import type { FC } from 'react';
import NextLink from 'next/link';
import { format } from 'date-fns';
import {
  Card,
  CardHeader,
  Divider,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TablePagination,
  TableRow
} from '@mui/material';
// import { customerApi } from '../../../api/customer-api';
import { ArrowRight as ArrowRightIcon } from '../../../icons/arrow-right';
// import type { CustomerInvoice } from '../../../types/customer';
import { MoreMenu } from '../../more-menu';
import { Scrollbar } from '../../scrollbar';
import { SeverityPill } from '../../severity-pill';

export const CustomerInvoices: FC = (props) => {

  return (
    <Card {...props}>
      <CardHeader
        action={<MoreMenu />}
        title="Recent Invoices"
      />
      <Divider />
      <Scrollbar>
        <Table sx={{ minWidth: 600 }}>
          <TableHead>
            <TableRow>
              <TableCell>
                ID
              </TableCell>
              <TableCell>
                Date
              </TableCell>
              <TableCell>
                Total
              </TableCell>
              <TableCell>
                Status
              </TableCell>
              <TableCell align="right">
                Actions
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
          </TableBody>
        </Table>
      </Scrollbar>
    </Card>
  );
};
