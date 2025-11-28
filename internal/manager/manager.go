package manager

import (
	"fmt"

	"github.com/codecrafters-io/kafka-starter-go/internal/metadata"
	"github.com/google/uuid"
)

type Manager struct {
	topicsByName map[string]*Topic
	topicsByUUID map[uuid.UUID]*Topic
}

var manager *Manager

func GetManager() (*Manager, error) {
	if manager == nil {
		return nil, fmt.Errorf("manager is not initialized")
	}

	return manager, nil
}

func InitManger(metadata *metadata.Metadata) (*Manager, error) {
	manager = &Manager{
		topicsByUUID: map[uuid.UUID]*Topic{},
		topicsByName: map[string]*Topic{},
	}
	for _, entry := range metadata.Topics {
		topicUUID := uuid.UUID(entry.TopicUUID.Value)
		topic := Topic{
			Name:       entry.Name.Value,
			UUID:       topicUUID,
			Partitions: []Partition{},
		}
		manager.topicsByName[entry.Name.Value] = &topic
		manager.topicsByUUID[topicUUID] = &topic
	}
	for _, entry := range metadata.Partitions {
		topicUUID := uuid.UUID(entry.TopicUUID.Value)
		partition := Partition{
			PartitionID:      entry.PartitionID.Value,
			Replicas:         []uint32{},
			InSyncReplicas:   []uint32{},
			RemovingReplicas: []uint32{},
			AddingReplicas:   []uint32{},
			Leader:           entry.Leader.Value,
			LeaderEpoch:      entry.LeaderEpoch.Value,
			PartitionEpoch:   entry.PartitionEpoch.Value,
			Directories:      []uuid.UUID{},
		}

		for _, replica := range entry.Replicas.Items {
			partition.Replicas = append(partition.Replicas, replica.Value)
		}
		for _, replica := range entry.SyncReplicas.Items {
			partition.InSyncReplicas = append(partition.InSyncReplicas, replica.Value)
		}
		for _, replica := range entry.RemovingReplicas.Items {
			partition.RemovingReplicas = append(partition.RemovingReplicas, replica.Value)
		}
		for _, replica := range entry.AddingReplicas.Items {
			partition.AddingReplicas = append(partition.AddingReplicas, replica.Value)
		}
		for _, directory := range entry.Directories.Items {
			partition.Directories = append(partition.Directories, uuid.UUID(directory.Value))
		}

		manager.topicsByUUID[topicUUID].Partitions = append(manager.topicsByUUID[topicUUID].Partitions, partition)
	}

	return manager, nil
}

type Topic struct {
	Name       string
	UUID       uuid.UUID
	Partitions []Partition
}

type Partition struct {
	PartitionID      uint32
	Replicas         []uint32
	InSyncReplicas   []uint32
	RemovingReplicas []uint32
	AddingReplicas   []uint32
	Leader           uint32
	LeaderEpoch      uint32
	PartitionEpoch   uint32
	Directories      []uuid.UUID
}

func (m *Manager) GetTopicByName(name string) (*Topic, error) {
	topic, exist := m.topicsByName[name]
	if !exist {
		return nil, fmt.Errorf("Topic with name (%s) Not Found", name)
	}

	return topic, nil
}

func (m *Manager) GetTopicByUUID(uuid uuid.UUID) (*Topic, error) {
	topic, exist := m.topicsByUUID[uuid]
	if !exist {
		return nil, fmt.Errorf("Topic with uuid (%s) Not Found", uuid.String())
	}

	return topic, nil
}
