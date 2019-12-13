.PHONY: all clean sequence  job logic relation ap

RELATION_OUTPUT=relation
JOB_OUTPUT=job
LOGIC_OUTPUT=logic
SEQUENCE_OUTPUT=seqsrv
AP_OUTPUT=apsrv

all: clean sequence  job logic relation ap

ap:
	go build -o ap/bin/${AP_OUTPUT} ap/main.go

relation:
	go build -o relation/bin/${RELATION_OUTPUT} relation/main.go

job:
	go build -o job/bin/${JOB_OUTPUT}  job/main.go

logic:
	go build -o logic/bin/${LOGIC_OUTPUT} logic/main.go

sequence:
	go build -o sequence/bin/${SEQUENCE_OUTPUT} sequence/main.go



clean:
	rm -f ap/bin/${AP_OUTPUT}
	rm -f job/bin/${JOB_OUTPUT}
	rm -f logic/bin/${LOGIC_OUTPUT}
	rm -f sequence/bin/${SEQUENCE_OUTPUT}
	rm -f relation/bin/${RELATION_OUTPUT}