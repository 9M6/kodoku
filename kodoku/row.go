package kodoku

// Row returns the row of the tile at the given index.
type Row []Tile

// Sort sorts the row.
func (r Row) Len() int           { return len(r) }
func (r Row) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Row) Less(i, j int) bool { return r[i] < r[j] }
