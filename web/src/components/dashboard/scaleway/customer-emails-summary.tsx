import { useState, useCallback, useEffect } from 'react';
import type { FC } from 'react';
import { format } from 'date-fns';
import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TextField,
  Typography
} from '@mui/material';
// import { customerApi } from '../../../api/customer-api';
import { ArrowRight as ArrowRightIcon } from '../../../icons/arrow-right';
// import type { CustomerEmail } from '../../../types/customer';

const emailOptions = [
  'Resend last invoice',
  'Send password reset',
  'Send verification'
];

export const CustomerEmailsSummary: FC = (props) => {
  const [emailOption, setEmailOption] = useState<string>(emailOptions[0]);

  return (
    <Card {...props}>
      <CardHeader title="Emails" />
      <Divider />
      <CardContent>
        <TextField
          name="option"
          onChange={(event): void => setEmailOption(event.target.value)}
          select
          SelectProps={{ native: true }}
          sx={{
            width: 320,
            maxWidth: '100%'
          }}
          value={emailOption}
        >
          {emailOptions.map((option) => (
            <option
              key={option}
              value={option}
            >
              {option}
            </option>
          ))}
        </TextField>
        <Box sx={{ mt: 2 }}>
          <Button
            endIcon={<ArrowRightIcon fontSize="small" />}
            variant="contained"
          >
            Send email
          </Button>
        </Box>
      </CardContent>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>
              Mail Type
            </TableCell>
            <TableCell>
              Date
            </TableCell>
          </TableRow>
        </TableHead>
      </Table>
    </Card>
  );
};
