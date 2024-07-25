const os = @import("std").os;

pub const CPU = struct {
    usage: f64,
    perCoreUsage: []f64,
    loadAverage: []f64,
};

pub const DISK = struct {
    total: u64,
    used: u64,
    free: u64,
    readBytes: u64,
    writeBytes: u64,
    readOperations: u64,
    writeOperations: u64,
};

// TODO: Define structs

pub fn cpu() !CPU {
    // TODO: CPU usage

    // TODO: Per-core usage

    // TODO: Average usage

    var cpuMetrics = CPU{
        .usage = getCPUMetrics(),
        .perCoreUsage = getPerCoreUsage(),
        .loadAverage = getLoadAverage(),
    };
    return cpuMetrics;
}

// TODO: Define calls
