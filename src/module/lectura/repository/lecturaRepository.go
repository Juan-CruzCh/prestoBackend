package repository

import (
	"context"
	"errors"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/lectura/dto"
	"prestoBackend/src/module/lectura/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LecturaRepository interface {
	CrearLectura(lectura *model.Lectura, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarLectura(filter *dto.BuscadorLecturaDto, ctx context.Context) (*[]bson.M, error)
	ActualizarLectura(ctx context.Context)
	EliminarLectuta()
	NumeroDeLecturaPorMedidor(medidor *bson.ObjectID, ctx context.Context) (int, error)
	CantidadLecturas(ctx context.Context) (int, error)
	ContarLecturasPorMedidorYEstado(medidor *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (int, error)
	BuscarLecturaPorId(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*model.Lectura, error)
	ActualizarEstadoLectura(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*mongo.UpdateResult, error)
	UltimaLecturaMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Lectura, error)
	LecturasPorMedidor(medidor *bson.ObjectID, ctx context.Context) ([]model.Lectura, error)
	HistorialLecturaMedidor(medidor *bson.ObjectID, ctx context.Context) ([]model.Lectura, error)
}

type lecturaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewLecturaRepository(db *mongo.Database) LecturaRepository {
	return &lecturaRepository{
		db:         db,
		collection: db.Collection("Lectura"),
	}

}

func (r *lecturaRepository) CrearLectura(lectura *model.Lectura, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "medidor": lectura.Medidor, "mes": lectura.Mes, "gestion": lectura.Gestion})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, errors.New("la lectura ya se encuetra registrada")
	}
	resultado, err := r.collection.InsertOne(ctx, lectura)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *lecturaRepository) ListarLectura(filter *dto.BuscadorLecturaDto, ctx context.Context) (*[]bson.M, error) {
	f1, f2, err := utils.NormalizarRangoDeFechas(filter.FechaInicio, filter.FechaFin)
	if err != nil {

		return nil, err
	}
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "fecha", Value: bson.D{
					{Key: "$gte", Value: f1},
					{Key: "$lte", Value: f2},
				}},
			}},
		},

		utils.Lookup("Medidor", "medidor", "_id", "medidor"),
		utils.Unwind("$medidor", false),
	}

	pipeline = append(pipeline, bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "numeroMedidor", Value: "$medidor.numeroMedidor"},
			{Key: "gestion", Value: 1},
			{Key: "mes", Value: 1},
			{Key: "lecturaActual", Value: 1},
			{Key: "lecturaAnterior", Value: 1},
			{Key: "consumoTotal", Value: 1},
			{Key: "costoApagar", Value: 1},
			{Key: "estado", Value: 1},
			{Key: "_id", Value: 1},
		}},
	})
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []bson.M
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}
func (r *lecturaRepository) ActualizarLectura(ctx context.Context) {

}
func (r *lecturaRepository) EliminarLectuta() {

}

func (r *lecturaRepository) NumeroDeLecturaPorMedidor(medidor *bson.ObjectID, ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "medidor": medidor})
	if err != nil {
		return 0, err
	}
	cantidad += 1
	return int(cantidad), nil

}

func (r *lecturaRepository) CantidadLecturas(ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	cantidad += 1
	return int(cantidad), nil

}

func (r *lecturaRepository) ContarLecturasPorMedidorYEstado(medidor *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "estado": estado, "medidor": medidor})
	if err != nil {
		return 0, err
	}
	return int(cantidad), nil

}

func (r *lecturaRepository) BuscarLecturaPorId(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*model.Lectura, error) {

	var data model.Lectura
	err := r.collection.FindOne(ctx, bson.M{"flag": enum.FlagNuevo, "estado": estado, "_id": lectura}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *lecturaRepository) ActualizarEstadoLectura(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*mongo.UpdateResult, error) {
	var filter bson.M = bson.M{"_id": lectura, "flag": enum.FlagNuevo}
	var update bson.D = bson.D{{Key: "$set", Value: bson.D{{Key: "estado", Value: estado}}}}

	resultado, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *lecturaRepository) UltimaLecturaMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Lectura, error) {
	var lectura model.Lectura
	opts := options.FindOne().SetSort(bson.D{{Key: "numeroLectura", Value: -1}})

	err := r.collection.FindOne(ctx, bson.M{"medidor": medidor, "flag": enum.FlagNuevo}, opts).Decode(&lectura)
	if err != nil {
		return nil, err
	}
	return &lectura, nil
}

func (r *lecturaRepository) LecturasPorMedidor(medidor *bson.ObjectID, ctx context.Context) ([]model.Lectura, error) {
	var lecturas []model.Lectura = []model.Lectura{}

	cursor, err := r.collection.Find(ctx, bson.M{"medidor": medidor, "flag": enum.FlagNuevo, "estado": enum.LecturaPendiente})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &lecturas)
	if err != nil {
		return nil, err
	}
	return lecturas, nil
}

func (r *lecturaRepository) HistorialLecturaMedidor(medidor *bson.ObjectID, ctx context.Context) ([]model.Lectura, error) {
	var lecturas []model.Lectura = []model.Lectura{}

	cursor, err := r.collection.Find(ctx, bson.M{"medidor": medidor, "flag": enum.FlagNuevo, "gestion": "2025"})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &lecturas)
	if err != nil {
		return nil, err
	}
	return lecturas, nil
}
