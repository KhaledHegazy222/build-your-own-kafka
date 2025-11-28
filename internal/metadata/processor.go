package metadata

type Processor interface {
	Process() error
}

func (fr *FeatureLevelRecord) Process() error {
	cfg, err := GetMetadata()
	if err != nil {
		return err
	}
	cfg.Features = append(cfg.Features, *fr)

	return nil
}

func (tr *TopicRecord) Process() error {
	cfg, err := GetMetadata()
	if err != nil {
		return err
	}
	cfg.Topics = append(cfg.Topics, *tr)

	return nil
}

func (pr *PartitionRecord) Process() error {
	cfg, err := GetMetadata()
	if err != nil {
		return err
	}

	cfg.Partitions = append(cfg.Partitions, *pr)

	return nil
}
