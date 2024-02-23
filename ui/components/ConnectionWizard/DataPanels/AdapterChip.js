import { withStyles, Chip } from '@material-ui/core/';

const chipStyles = (theme) => ({
  chipIcon: { width: theme.spacing(2.5) },
  chip: { marginRight: theme.spacing(1), marginBottom: theme.spacing(1) },
});

const AdapterChip = withStyles(chipStyles)(
  ({ classes, handleClick, handleDelete, label, image, isActive }) => (
    <Chip
      label={label}
      onDelete={handleDelete}
      onClick={handleClick}
      icon={<img src={image} className={classes.chipIcon} />}
      className={classes.chip}
      key={label + '-key'}
      variant={isActive ? 'outlined' : 'default'}
    />
  ),
);

export default AdapterChip;
